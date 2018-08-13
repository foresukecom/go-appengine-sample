$(document).ready(function () {
    $.ajax({
        url: '/hoge',
        type: "get"
    }).done(
        function (data) {
            $('#content').append(data);
        });
});

