$(document).ready(function () {
    setTimeout(function () {
        $("#uiBar").animate(
            {width:"100%"},
            {duration: 8000,
                complete: function() {
                    $("#uiCounter").text("LOADED");
                    $("#uiBar").css("background","#00CCB1")
                },
                step:function(now, fx){
                    $("#uiCounter").text(Math.round(now) + " %");}
            }

        );
    }, 1000);
});