package api

import (
	"net/http"
	"os"

	routing "github.com/julienschmidt/httprouter"
)

// HomeLink implements GET / Just a test msg to see if it works
func HomeLink(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	homeHTML := `
	<style>
		.box{
			-moz-border-radius: 6px;
			-webkit-border-radius: 6px;
			background-color: #fbf8ff;
			background-image: url(../Images/icons/Pencil-48.png);
			background-position: 9px 0px;
			background-repeat: no-repeat;
			border: solid 1px #3498db;
			border-radius: 6px;
			line-height: 18px;
			overflow: hidden;
			padding: 15px 60px;
			width: 300px;
			margin: auto;
				margin-top: auto;
			box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em, rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em, rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
			margin-top: 200px;
			text-align:center;
		}
	</style>
	<div class="box">
		Salam Goloooo, It works!
		<p>
			Navigate to the <a href="/ui/" >Web UI</a>
		</p>
	</div>
	<script>
		setTimeout( () => {window.location.href="/ui/"}, 500)
	</script>
	`
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp.Write([]byte(homeHTML))
}

/*-------------------------*/

func UI(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	rootPath := os.Getenv("EXEC_PATH")
	if rootPath == "" {
		rootPath = "./"
	}

	http.FileServer(http.Dir(rootPath)).ServeHTTP(resp, req)
}

/*-------------------------*/
