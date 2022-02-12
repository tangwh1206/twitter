window.onload = initPage();

function getCookie(name) {
    var strcookie = document.cookie
    var arrcookie = strcookie.split("; ");
    for ( var i = 0; i < arrcookie.length; i++) {
        var arr = arrcookie[i].split("=");
        if (arr[0] == name)
            return arr[1];
    }
    return "";
}

function logged() {
    return getCookie("twitter_session_v1").length > 0 ? true : false;
}

function initPage() {
    if (logged()) {
        var navItemUser = document.getElementById("user");
        navItemUser.innerHTML = "";
        navItemUser.href = "#todo-user-home-page"
        navItemUser.text = "已登录"

        var userDropdownContent = document.getElementById("user-dropdown-content");
        ch1 = document.createElement("a");
        ch1.text = "个人主页"
        ch1.href = "#todo-user-home-page"

        ch2 = document.createElement("a");
        ch2.text = "退出登录"
        ch2.href = "/tiwtter/user/logout"

        userDropdownContent.appendChild(ch1)
        userDropdownContent.appendChild(ch2)
    } else {
        alert("未登录");
    }
}