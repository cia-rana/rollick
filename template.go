package main

var index = `
<!DOCTYPE html>
<html>
<head>
	<script src="https://code.jquery.com/jquery-1.7.0.min.js" integrity="sha256-/05Jde9AMAT4/o5ZAI23rUf1SxDYTHLrkOco0eyRV84=" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/echo-js@1.7.3/src/echo.min.js"></script>
</head>
<body>
<script>
window.onload = function() {
	echo.init();

	let images = document.getElementsByClassName("image")
	let imageLen = images.length;

	if (imageLen == 0) {
		return;
	}
	
	for (let i = 1; i < imageLen; i++) {
		images[i].style.display = "none";
	}

	let pos = 0;
	document.body.onkeydown = function(e){
		switch (e.code) {
		case "ArrowLeft":
			images[pos].style.display = "none";
			pos = (pos - 1 + imageLen) % imageLen;
			images[pos].style.display = "block";
			break;
		case "ArrowRight":
			images[pos].style.display = "none";
			pos = (pos + 1) % imageLen;
			images[pos].style.display = "block";
			break;
		}
	}
}
</script>
<div class="viewer">
	{{range .}}
	<div class="image">
		<img data-echo="./{{.}}">
	</div>
	{{end}}
</div>
</body>
</html>
`
