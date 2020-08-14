package permissions

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	"gopackage/loger"
)

var Casbin *casbin.Enforcer

func Setup(path string, db *gorm.DB){

	mysql := gormadapter.NewAdapterByDB(db)
	Casbin = casbin.NewEnforcer(path, mysql)

	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.

	// Load the policy from DB.
	err := Casbin.LoadPolicy()

	if err != nil {
		loger.Error("# Casbin SetUp Failed.", err)
		return
	}

	loger.Info("Casbin SetUp Success...")
}


func TestRBAC(){

	/**
	p策略 -> 用户/角色-资源-访问的映射关系
	p, data2_admin, data2, read
	由于data2_admin可以表示为用户或角色， 映射关系即可以为： 用户data2_admin对资源data2拥有read的权限 或 角色data2_admin对资源data2拥有read的权限。
	RBAC通常使用第二种方式表示，角色-资源映射关系，然后结合g策略

	g策略 -> 用户/资源/角色 - 角色的映射关系
	g, alice, data2_admin
	alice可以是用户、资源或角色的其中一种， Cabin 只是将其识别为一个字符串。
	alice是角色 data2_admin的一个成员。重点关注其中两种表示含义即可：用户alice属于角色data2_admin， 子角色alice属于角色data2_admin

	p	alice1	/user/login	read
	p	alice2	/user/login	read
	p	role1	/user/login	read
	g	user1	role1
	g	user2	role1
	g	user1	role2

	*/
	loger.Info(Casbin.GetAllRoles())		//所有角色 [role1 role2]
	loger.Info(Casbin.GetAllSubjects())	//所有策略对象，即包含资源/用户/角色 [alice1 alice2 role1]
	loger.Info(Casbin.GetAllActions())		//所有操作 [read]
	loger.Info(Casbin.GetAllObjects())		//所有访问对象（访问资源） /user/login
	loger.Info(Casbin.GetUsersForRole("role1"))	//获取属于角色role1的用户[user1 user2] <nil>
	loger.Info(Casbin.GetRolesForUser("user1")) //获取用户user1拥有的角色 [role1 role2] <nil>]
	loger.Info(Casbin.GetRolesForUser("user2")) //获取用户user2拥有的角色[role1] <nil>

	sub := "role1" 			// 用户或角色
	obj := "/user/login" 	// 将要访问的资源
	act := "read" 			// 对资源执行的操作
	loger.Info("casbin Enforce: ", Casbin.Enforce(sub, obj, act))	// [casbin Enforce：  true]

	sub = "alice1" 			// 用户或角色
	obj = "/user/login" 	// 将要访问的资源
	act = "read" 			// 对资源执行的操作
	loger.Info("casbin Enforce: ", Casbin.Enforce(sub, obj, act))	// [casbin Enforce：  true]

	sub = "alice1" 			// 用户或角色
	obj = "/user/logout" 	// 将要访问的资源
	act = "read" 			// 对资源执行的操作
	loger.Info("casbin Enforce: ", Casbin.Enforce(sub, obj, act))	// [casbin Enforce：  false]


}