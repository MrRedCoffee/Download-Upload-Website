package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// uploadFile aceita o arquivo e salva no diretorio atual
func uploadFile(w http.ResponseWriter, r *http.Request) {
	// a função FormFile pega o id de entrada do arquivo por metodo POST
	file, header, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// cria o arquivo na pasta raiz
	out, err := os.Create("downloads/temp/" + header.Filename)
	if err != nil {
		fmt.Println("Não foi possivel criar o arquivo no servidor.")
		fmt.Println(err)
		return
	}

	// escreve no arquivo criado usando metodo POST. write the content from POST to the file
	io.Copy(out, file)
	file.Close()
	out.Close()

	fmt.Fprintf(w, "Arquivo enviado com sucesso: ")
	fmt.Fprintf(w, header.Filename)
}


func upServer() {
  	//define a porta
	port := ":80"
	fmt.Println("Rodando server HTTP na porta", port)

  	//define diretorio raiz
	root_dir := http.FileServer(http.Dir("./static"))

	download_dir := http.FileServer(http.Dir("./downloads"))
  	http.Handle("/", root_dir)
  	http.Handle("/download/", http.StripPrefix("/download/", download_dir))

  	//define o diretorio que realiza a função upload
  	//o /upload é chamado no action do html
  	http.HandleFunc("/upload", uploadFile)

  	//inicia o server
  	http.ListenAndServe(port, nil)
}

func main() {
	upServer()
}