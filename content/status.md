# Service Status

<script type="text/lodash-template" id="status-template">
<hr>
<div class="check row d-flex align-items-center">
    <div class="state col-2 <%= StateClass %>">
        <span class="badge status-badge <%= BadgeClass %>"><%= State %></span>
    </div>
    <div class="service col">
        <b><%= Service %></b><br>
        <a href="https://updown.io/<%= Token %>"><small><%= LastCheck %></small></a>
    </div>
    <div class="uptime col-3 d-none d-sm-block">
        <span class="badge status-badge <%= UptimeBadgeClass %>"><%= Uptime %> %</span>
    </div>
</div>
</script>

<div id="status-container">
<div class="check row d-flex align-items-center">
    <div class="state col-2">State</div>
    <div class="service col">Service</div>
    <div class="uptime col-3 d-none d-sm-block">30d Uptime</div>
</div>
</div>