(function () {
    var tpl = $('#status-template')
    if (!tpl) return
    var ctr = $('#status-container')
    if (!ctr) return
    var tplFn = _.template(tpl.text())

    $.getJSON('/.netlify/functions/status', function (ss) {
        $.each(ss, function (_, s) {
            s.State = s.OK ? 'Up' : 'Down'
            s.StateClass = s.OK ? 'state-up' : 'state-down'
            s.BadgeClass = s.OK ? 'badge-success' : 'badge-danger'
            s.UptimeBadgeClass = s.Uptime >= 99.9 ? 'badge-success' : 'badge-warning'
            ctr.append(tplFn(s))
        })
    })
})()