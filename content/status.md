# Status

<script type="text/lodash-template" id="status-template">
<div class="check row d-flex align-items-center">
    <div class="state col-2 <%= StateClass %>"><span class="badge"><%= State %></span></div>
    <div class="service col"><%= Service %></div>
    <div class="uptime col-3 d-none d-sm-block"><%= Uptime %> %</div>
</div>
</script>

<div id="status-container">
</div>