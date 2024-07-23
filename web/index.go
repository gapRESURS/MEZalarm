package web

import (
	"MEZ/dataset"
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `
<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Оповещение МЭЗ</title>
	<link rel="shortcut icon" type="image/png" href="https://www.gapresurs.ru/favicon.ico">
    <link href="CSS/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
	<link href="CSS/resurs.css" rel="stylesheet" >
</head>

<body class="bg-dark text-white">

		<nav class="navbar navbar-expand-lg navbar-light blagoyar-green">
			<div class="container-fluid">
				<a class="navbar-brand" onclick="toggleTheme()">
					<img src="/IMG/logo.svg" height="30" hspace="10">
				</a>
			</div>
		</nav>

<div class="mx-2 my-3">
	<table class="table table-dark table-striped table-hover">
		<thead>
			<tr>
			  <th scope="col">МЭЗ</th>
			  <th scope="col">Оповещение</th>
			  <th scope="col">Температура</th>
			  <th scope="col">Влажность</th>
			</tr>
		</thead>
		<tbody>
			<tr id="NMEZ">
			  <th scope="row"><div id="ping" data-ip="%s">Невинномысский</div></th>
			  <td id="relay">-</td>
			  <td id="temperature">0°C</td>
			  <td id="humidity">0%%</td>
			</tr>
			<tr id="AMEZ">
			  <th scope="row"><div id="ping" data-ip="%s">Армавирский</div></th>
			  <td id="relay">-</td>
			  <td id="temperature">0°C</td>
			  <td id="humidity">0%%</td>
			</tr>
			<tr id="GMEZ">
			  <th scope="row"><div id="ping" data-ip="%s">Георгиевский</div></th>
			  <td id="relay">-</td>
			  <td id="temperature">0°C</td>
			  <td id="humidity">0%%</td>
			</tr>
			<tr id="SMEZ">
			  <th scope="row"><div id="ping" data-ip="%s">Сальский</div></th>
			  <td id="relay">-</td>
			  <td id="temperature">0°C</td>
			  <td id="humidity">0%%</td>
			</tr>
			<tr id="UMEZ">
			  <th scope="row"><div id="ping" data-ip="%s">Усть-Лабинский</div></th>
			  <td id="relay">-</td>
			  <td id="temperature">0°C</td>
			  <td id="humidity">0%%</td>
			</tr>
		</tbody>
	</table>

</div>
<div class="text-secondary fs-6 position-fixed bottom-0 start-50 translate-middle-x">
	К.А. Рыжков 2024
</div>

<script src="JS/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
<script src="JS/statusERD.js"></script>
<script src="JS/status.js"></script>

</body>
</html>
`, dataset.NMEZ_IP, dataset.AMEZ_IP, dataset.GMEZ_IP, dataset.SMEZ_IP, dataset.UMEZ_IP)

}
