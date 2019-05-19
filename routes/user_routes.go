package routes

import (
    "knn_contest/config"
    "knn_contest/sessions"
    
    "net/http"
    "github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
    println("post/signup")
    username := ctx.PostForm("username")
    email := ctx.PostForm("emailaddress")
    password := ctx.PostForm("password")
    passwordConf := ctx.PostForm("passwordconfirmation")

if password != passwordConf {
        println("Error: password and passwordConf does not match")
        ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
        return    
    }

    db := config.DummyDB()
    if err := db.SaveUser(username, email, password); err != nil {
        println("Error: "+ err.Error())
        ctx.Redirect(http.StatusSeeOther, "/")
        return
    }
    
    println("Signup success!!")
    println("  username: " + username)
    println("  email: " + email)
    println("  password: " + password)            

    ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")	
}

func UserLogIn(ctx *gin.Context) {
    println("post/login")
    username := ctx.PostForm("username")
    password := ctx.PostForm("password")

    db := config.DummyDB()
    user, err := db.GetUser(username, password)
    if err != nil {
        println("Error: " + err.Error())
        ctx.Redirect(http.StatusSeeOther, "/")
	return
    }

    println("Authentication Success!!")
    println("  username: " + user.Username)
    println("  email: " + user.Email)
    println("  password: " + user.Password)
    session := sessions.GetDefaultSession(ctx)
    session.Set("user", user)
    session.Save()
    user.Authenticate()

    println("Session saved.")
    println("  sessionID: " + session.ID)
    ctx.Redirect(http.StatusSeeOther, "/")       

    ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

func UserLogOut(ctx *gin.Context) {
    session := sessions.GetDefaultSession(ctx)
    session.Terminate()
    ctx.Redirect(http.StatusSeeOther, "/")
}