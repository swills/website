# Status

<div id="status-container">
Loading current service status...
</div>
<script>
function reqListener () {
    document.getElementById('status-container').innerHTML = this.responseText;
}
var r = new XMLHttpRequest();
r.addEventListener("load", reqListener);
r.open("GET", "/.netlify/functions/status");
r.send();
</script>