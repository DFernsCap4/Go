package main

import (
    "errors"
    "fmt"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type User struct {
    ID    int
    Name  string
    Email string
}

var users []User
var nextID = 1

func addUser(name, email string) {
    user := User{ID: nextID, Name: name, Email: email}
    users = append(users, user)
    nextID++
}

func getUser(id int) (User, error) {
    for _, user := range users {
        if user.ID == id {
            return user, nil
        }
    }
    return User{}, errors.New("user not found")
}

func updateUser(id int, name, email string) error {
    for i, user := range users {
        if user.ID == id {
            users[i].Name = name
            users[i].Email = email
            return nil
        }
    }
    return errors.New("user not found")
}

func deleteUser(id int) error {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return errors.New("user not found")
}

func main() {
    a := app.New()
    w := a.NewWindow("User Management")

    nameEntry := widget.NewEntry()
    emailEntry := widget.NewEntry()

    // Adding User Section
    addButton := widget.NewButton("Add User", func() {
        name := nameEntry.Text
        email := emailEntry.Text
        addUser(name, email)
        widget.ShowPopUpMessage("User added successfully", w)
    })

    // Getting User Section
    idEntry := widget.NewEntry()
    getUserButton := widget.NewButton("Get User", func() {
        id := idEntry.Text
        userID := 0
        fmt.Sscan(id, &userID)
        user, err := getUser(userID)
        if err != nil {
            widget.ShowPopUpMessage("Error: "+err.Error(), w)
        } else {
            widget.ShowPopUpMessage(fmt.Sprintf("User: %+v", user), w)
        }
    })

    // Updating User Section
    updateButton := widget.NewButton("Update User", func() {
        id := idEntry.Text
        userID := 0
        fmt.Sscan(id, &userID)
        name := nameEntry.Text
        email := emailEntry.Text
        err := updateUser(userID, name, email)
        if err != nil {
            widget.ShowPopUpMessage("Error: "+err.Error(), w)
        } else {
            widget.ShowPopUpMessage("User updated successfully", w)
        }
    })

    // Deleting User Section
    deleteButton := widget.NewButton("Delete User", func() {
        id := idEntry.Text
        userID := 0
        fmt.Sscan(id, &userID)
        err := deleteUser(userID)
        if err != nil {
            widget.ShowPopUpMessage("Error: "+err.Error(), w)
        } else {
            widget.ShowPopUpMessage("User deleted successfully", w)
        }
    })

    // Layout
    content := container.NewVBox(
        widget.NewLabel("Add or Update User"),
        widget.NewLabel("Name:"),
        nameEntry,
        widget.NewLabel("Email:"),
        emailEntry,
        addButton,
        widget.NewLabel("User ID:"),
        idEntry,
        getUserButton,
        updateButton,
        deleteButton,
    )

    w.SetContent(content)
    w.ShowAndRun()
}
