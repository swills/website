# Status

<script type="text/lodash-template" id="status-template">
<hr>
<div class="check row d-flex align-items-center">
    <div class="state col-2 <%= StateClass %>">
        <span class="badge <%= BadgeClass %>"><%= State %></span>
    </div>
    <div class="service col">
        <%= Service %><br>
        <small><%= LastCheck %></small>
    </div>
    <div class="uptime col-3 d-none d-sm-block"><%= Uptime %> %</div>
</div>
</script>

<div id="status-container">
<div class="check row d-flex align-items-center">
    <div class="col-2">State</div>
    <div class="col">Service</div>
    <div class="col-3">30d Uptime</div>
</div>
</div>