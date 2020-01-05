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
            ctr.append(tplFn(s))
        })
    })
})()