package em

import (
	"net/http"

	"github.com/ear7h/ear7h-net/api"
)

func init() {
	api.HandleFunc("/em", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
<!doctype html>
<body>
	<p>this isn't the most extraordinary piece of code</p><br>
	<p style="display: inline">but click </p>
	<button style="display: inline" onclick="goof()">here</button>
	<p style="display: inline">repeatedly</p>

	<div id="a-div">
	</div>

	<script>
		function goof() {
			let d = document.getElementById("a-div")

			let c = document.createElement("p")
			c.style.color = "#" + ((1 << 24) * Math.random() | 0).toString(16)
			c.innerHTML = "i can't wait to see you, goof"

			d.appendChild(c)
		}
	</script>
</body>
		`))
	})
}
