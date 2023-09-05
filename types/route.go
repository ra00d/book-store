package types

// package types
//
// import (
// 	"fmt"
// 	"log"
// 	"strings"
//
// 	"github.com/gofiber/fiber/v2"
// )
//
// type (
// 	RouterGroup struct {
// 		Name            string
// 		IsPublic        bool
// 		UserRole        string
// 		UserPermissions []string
// 	}
// )
//
// func CreateNewRouterGroup(app *fiber.App,
// 	config RouterGroup,
// 	handlers ...fiber.Handler,
// ) fiber.Router {
// 	if len(config.Name) == 0 {
// 		log.Fatal("The group Name is used as end point and can not be empty")
// 	}
//
// 	if len(config.UserRole) == 0 && len(config.UserPermissions) > 0 {
// 		log.Fatal("PERMISSIONS ARE NOT ALLOWED WITHOUT A ROLE ")
// 	}
// 	var group fiber.Router
// 	var groupHandlers []fiber.Handler
// 	if !config.IsPublic && len(config.UserRole) == 0 {
// 		group = app.Group(strings.ToLower(config.Name), handlers...)
// 	} else if len(config.UserRole) > 0 && len(config.UserPermissions) == 0 {
// 		groupHandlers = append(groupHandlers, auth.WithRole(strings.Title(config.UserRole)))
// 		groupHandlers = append(groupHandlers, handlers...)
// 		fmt.Println("in user no PERMISSIONS")
// 		group = app.Group(strings.ToLower(config.Name), groupHandlers...)
// 	} else if len(config.UserRole) > 0 && len(config.UserPermissions) > 0 {
// 		groupHandlers = append(groupHandlers, auth.WithRole(strings.ToLower(config.UserRole)),
// 			// TODO FIX THIS TO ALLOW ARRAY OF PERMISSIONS
// 			auth.WithPermission(config.UserPermissions[0]))
// 		groupHandlers = append(groupHandlers, handlers...)
// 		group = app.Group(strings.ToLower(config.Name), groupHandlers...)
// 	} else {
// 		group = app.Group(strings.ToLower(config.Name), handlers...)
// 	}
// 	return group
// }
