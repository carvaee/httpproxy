package HTTPServer

import (
	"entry_task/RPCServer"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const ImgDir string = "./data"
func init() {
	dir, err := os.Stat(ImgDir)
	if err != nil {
		fmt.Println("[info] the directory of img is not exist, try to create")
		_ = os.Mkdir(ImgDir, os.ModePerm)
	}  else {
		if !dir.IsDir() {
			err = os.Remove(ImgDir)
			os.Mkdir(ImgDir, os.ModePerm)
		}
	}
	dir, err = os.Stat(ImgDir)
	if err !=nil || !dir.IsDir() {
		fmt.Println("[error] dir =",dir.IsDir(), ", error: ",err.Error())
		panic("image dir is not exit and can't be created")
	}
}
func StartHTTPServer () {
	http.HandleFunc("/", index)
	http.HandleFunc("/profile", profile)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./HTTPServer/template/index.html")
	if err != nil {
		fmt.Printf("[error] Parse template <index> error, err: %v", err)
		return
	}
	err = t.Execute(w,nil)
	if err != nil {
		fmt.Printf("[error] Render template <index> error, err: %v", err)
		return
	}
}
func profile(w http.ResponseWriter, req *http.Request) {
	postFlag := req.PostFormValue("postFlag")
	switch postFlag {
	case "login":
		// 验证用户名 密码
		userName := req.PostFormValue("admin")
		password := req.PostFormValue("password")
		user, err := RPCServer.Login(userName, password)
		if err != nil {
			return
		}
		t, err := template.ParseFiles("./HTTPServer/template/profile.html")
		if err != nil {
			fmt.Printf("[error] Parse template <profile> error, err: %v", err)
			return
		}
		err = t.Execute(w, user)
		//fmt.Printf("%v", *user)
		if err != nil {
			fmt.Printf("[error] Render template <profile> error, err: %v", err)
			return
		}

	case "ModifyNickName":

	case "ModifyImage":
		userName:=req.PostFormValue("admin")
		img, handle, _ :=req.FormFile("image")
		temp := strings.Split(handle.Filename, ".")
		extension := strings.ToLower(temp[len(temp)-1])
		fileName := fmt.Sprintf("%s.%s", userName, extension)
		fmt.Printf("filename=%s\n", fileName)
		saveImg(img, fileName)
	}
}

func saveImg(img multipart.File,imgName string){
	fullPath := ImgDir + "/" + imgName
	file,err :=os.Create(fullPath)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	_,err =io.Copy(file,img)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	_= file.Close()
}