<template>
<div class="container">
    <h5>Thank to:</h5>
    <a class="body-1" href="#" @click="open('https://golang.org/')" >Go</a>
    <br>
    <a class="body-1" href="#" @click="open('https://github.com/gin-gonic/gin')" >Gin-gonic (Go Package)</a>
    <br>
    <a class="body-1" href="#" @click="open('https://github.com/zserge/webview')" >Webview (Go Package)</a>
    <br>
    <a class="body-1" href="#" @click="open('https://github.com/phayes/freeport')" >Freeport (Go Package)</a>
    <br>
    <a class="body-1" href="#" @click="open('https://vuejs.org')" >Vue</a>
    <br>
    <a class="body-1" href="#" @click="open('https://vuetifyjs.com')" >Vuetify</a>
    <br>
    <a class="body-1" href="#" @click="open('https://github.com/axios/axios')" >Axios</a>
    <br>
    <a class="body-1" href="#" @click="open('https://github.com/FranckFreiburger/http-vue-loader')" >httpVueLoader</a>
</div>
</template>
<script>
module.exports = {
    data: function() {
        return {
        }
    },
    mounted: function() {
    },
    methods: {
        open: function(url){
            app.api({
                method:"GET",
                url:"/api/systool/url/open",
                params:{url:url}
            }).then(data => {
                console.log("api quit - data")
                console.log(data)
            })
            .catch(error => {
                console.log("api quit - ERROR")
                console.log(error)
            });
        }
    }
}
</script>
<style>
</style>
