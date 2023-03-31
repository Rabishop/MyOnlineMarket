function uploadGame() {
    var gameName = document.forms["uploadGameTable"]["gameName"].value;
    var gamePrice = Number(document.forms["uploadGameTable"]["gamePrice"].value);
    var gameInfo = document.forms["uploadGameTable"]["gameInfo"].value;
    var gameImg = document.getElementById("gameImg").src;
    var gameGroup = document.forms["uploadGameTable"]["gameGroup"];
    var gameType = ""

    for (var i = 0; i < gameGroup.length; i++) {
        if (gameGroup[i].checked == true) {
            gameType += gameGroup[i].value + ";"
        }
    }

    var settings = {
        "url": "http://localhost:8080/game/upload",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "gamePrice": gamePrice,
            "gameName": gameName,
            "gameInfo": gameInfo,
            "gameImg": gameImg,
            "gameType": gameType
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            window.location.href = 'result.html';
        } else {
            alert("Fail!");
        }

    });

}

function getGameIndex() {
    var settings = {
        "url": "http://localhost:8080/game/index",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            // console.log(response);

            // console.log(response["gameIndex"][1]["gameItem"][1])
            for (var i = 0; i < response["gameIndex"].length; i++) {

                //创建最外层div.most-popular
                var divType = $('<div/>', {
                    class: 'most-popular'
                });

                //创建div.row
                var divRow = $('<div/>', {
                    class: 'row',
                    id: 'row' + i
                });

                //创建div.col-lg-12
                var divColLg12 = $('<div/>', {
                    class: 'col-lg-12'
                });

                //创建div.heading-section
                var divHeadingSection = $('<div/>', {
                    class: 'heading-section'
                });

                //创建h4标签
                var h4 = $('<h4/>', {
                    html: '<em>' + response["gameType"][i]["typeName"] + '</em>'
                });

                //将h4添加到div.heading-section中
                divHeadingSection.append(h4);

                //创建div.row#row
                var divRowRow = $('<div/>', {
                    class: 'row',
                    id: 'row'
                });

                //创建div.col-lg-12
                var divColLg12Two = $('<div/>', {
                    class: 'col-lg-12'
                });

                //创建div.main-button
                var divMainButton = $('<div/>', {
                    class: 'main-button'
                });

                //创建a标签
                var a = $('<a/>', {
                    href: 'browse.html',
                    html: 'Discover More'
                });

                //将a添加到div.main-button中
                divMainButton.append(a);

                //将所有元素层层添加
                divColLg12Two.append(divMainButton);
                divRowRow.append(divColLg12Two);
                divColLg12.append(divHeadingSection, divRowRow);
                divRow.append(divColLg12);
                divType.append(divRow);

                //将最外层div.most-popular添加到body中


                $('#container').append(divType);

                for (var j = 0; j < response["gameIndex"][i]["gameItem"].length; j++) {

                    // 创建一个 div 元素，并添加相应的类
                    var div = $('<div>').addClass('col-lg-3 col-sm-6');

                    // 创建 item 元素
                    var item = $('<div>').addClass('item');
                    div.append(item);

                    // 创建图片元素并设置属性
                    var img = $('<img>').attr('src', response["gameIndex"][i]["gameItem"][j]["gameImg"]).attr('style', "height: 216px; width: 216px;").attr('alt', '');
                    item.append(img);

                    // 创建标题 h4 元素和 span 元素，并设置相应的文本
                    var h4 = $('<h4>').text(response["gameIndex"][i]["gameItem"][j]["gameName"]);
                    var span = $('<span>').text(response["gameIndex"][i]["gameItem"][j]["gameUploader"]);
                    h4.append('<br>').append(span);
                    item.append(h4);

                    // 创建 ul 元素，里面添加两个 li 元素
                    var ul = $('<ul>');
                    var li1 = $('<li>').append($('<i>').addClass('fa fa-star')).append(response["gameIndex"][i]["gameItem"][j]["gameStar"]);
                    var li2 = $('<li>').append($('<i>').addClass('fa fa-usd')).append(response["gameIndex"][i]["gameItem"][j]["gamePrice"]);
                    ul.append(li1).append(li2);
                    item.append(ul);

                    // 将创建的 div 元素添加到页面中
                    $('#row' + i).append(div);

                    // console.log(response["gameIndex"][i]["gameItem"][j]);
                }

                $('#row' + i).append(divColLg12Two);
            }

        } else {
            alert("Get Index Fail!");
        }
    });

}
