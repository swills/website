(function () {
    var tpl = $('#status-template')
    if (!tpl) return
    var ctr = $('#status-container')
    if (!ctr) return
    var tplFn = _.template(tpl.text())

    $.getJSON('/.netlify/functions/status', function (ss) {
        $.each(ss, function (s) {
            ctr.append(tplFn(s))
        })
    })
})()