(function ($) {
    function Slide(ele, options) {
        this.$ele = $(ele)
        this.options = $.extend({
            speed: 1000,
            delay: 3000
        }, options)
        this.states = [
            { '&zIndex': 3, width: 98, height: 180, top: 71, left: 130, $opacity:  1 },
            { '&zIndex': 3, width: 130, height: 238, top: 71, left: 262, $opacity: 1 },
            { '&zIndex': 3, width: 100, height: 183, top: 71, left: 394, $opacity: 1 },
            { '&zIndex': 3, width: 98, height: 180, top: 71, left: 526, $opacity: 1},
        ]
        this.lis = this.$ele.find('li')
        this.interval
        this.$ele.find('section:nth-child(2)').on('click', function () {
            this.stop()
            this.next()
            this.play()
        }.bind(this))
        // 点击切换到上一张
        this.move()
        this.play()
    }


    Slide.prototype = {
        move: function () {
        this.lis.each(function (i, el) {
                $(el)
                    .css('z-index', this.states[i]['&zIndex'])
                    .finish().animate(this.states[i], this.options.speed)
                    // .stop(true,true).animate(states[i], 1000)
                    .find('img').css('opacity', this.states[i].$opacity)
            }.bind(this))
        },
        // 让轮播图切换到下一张
        next: function () {

            this.states.unshift(this.states.pop())
            this.move()
        },
        // 让轮播图滚动到上一张
        prev: function () {

            this.states.push(this.states.shift())
            this.move()
        },
        play: function () {

            this.interval = setInterval(function () {//这个this指window
                // setInterval、setTimeOut 中的this指向window

                // states.unshift(states.pop())       //从后往前走
                // states.push(states.shift())     //从前往后走
                this.next()
            }.bind(this), this.options.delay)
        },
        // 停止自动播放
        stop: function () {
            // var _this = this
            clearInterval(this.interval)
        }

    }
    $.fn.zySlide = function (options) {
        this.each(function (i, ele) {
            new Slide(ele, options)
        })
        return this
    }
})(jQuery)