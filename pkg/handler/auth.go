package handler

import (
	"net/http"

	"github.com/Manzo48/todo-app"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)


func (h *Handler) signUp(c *gin.Context) {
    logrus.Println("Handling sign-up request")

    var input todo.User
    if err := c.BindJSON(&input); err != nil {
        logrus.Printf("BindJSON error: %v", err)
        newErrorResponse(c, http.StatusBadRequest, "invalid input body")
        return
    }

    logrus.Printf("Creating user: %+v", input)
    id, err := h.services.Authorization.CreateUser(input)
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
            logrus.Printf("Username conflict: %v", err)
            newErrorResponse(c, http.StatusConflict, "username already exists")
            return
        }
        logrus.Printf("CreateUser error: %v", err)
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    logrus.Printf("User created with ID: %d", id)

    // Добавим дополнительное логирование перед отправкой ответа
    response := gin.H{
        "id": id,
        }
    logrus.Printf("Sending response: %+v", response)

    // Отправим JSON-ответ
    c.JSON(http.StatusOK, response); 
    logrus.Println("Response sent successfully")
}

func (h *Handler) signIn(c *gin.Context){

}

func (h *Handler) getAllUsers(c *gin.Context) {
    logrus.Println("Handling getAllUsers request")

    users, err := h.services.Authorization.GetAllUsers()
    if err != nil {
        logrus.Printf("GetAllUsers error: %v", err)
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    logrus.Printf("Users retrieved: %+v", users)
    c.JSON(http.StatusOK, users)
}