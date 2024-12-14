package handlers

import (
	"fmt"
	"net/http"
)

func GetBaseRoute(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "text/html")

	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Rotas Disponíveis</title>
		<style>
			body { font-family: Arial, sans-serif; line-height: 1.6; margin: 2em; }
			a { text-decoration: none; color: #007BFF; }
			a:hover { text-decoration: underline; }
			.container { max-width: 600px; margin: 0 auto; }
			h1 { color: #333; }
			ul { padding: 0; list-style: none; }
			li { margin: 0.5em 0; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Rotas Disponíveis</h1>
			<ul>
				<li><a href="/">Home - Rotas Disponíveis</a></li>
				<li><a href="/people">/people - Lista de Pessoas</a></li>
			</ul>
		</div>
	</body>
	</html>
	`

	fmt.Fprint(responseWriter, html)
}
