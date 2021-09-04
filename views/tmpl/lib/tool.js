var Tool = (function () {
    var paramsCache;

    return {
        params: function (key) {
            if (paramsCache === undefined) {
                paramsCache = {}
                window.location.search.replace("?", "").split("&").forEach((pair) => {
                    kv = pair.split("=")
                    paramsCache[kv[0]] = kv[1]
                })
            }
            return paramsCache[key]
        },

        todayStr: function () {
            var tod = new Date()
            var mon = tod.getMonth() + 1, mon = mon < 10 && ("0" + mon) || mon
            var date = tod.getDate() < 10 && ("0" + tod.getDate()) || tod.getDate()
            return "" + tod.getFullYear() + "-" + mon + "-" + date
        },

        dateAfterMonths: function (dateStr, n) {
            var ymd = dateStr.split("-"), y = ymd[0], m = ymd[1], d = ymd[2]
            m = parseInt(m) + n
            var addYear = parseInt(m / 12)
            if (addYear > 0) {
                y = parseInt(y) + addYear
                m = m % 12
            }
            m = m < 10 && ("0" + m) || m
            return "" + y + "-" + m + "-" + d
        }
    }
}())