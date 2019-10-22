<template>
<div class="container">
    <h3>home</h3>
    <v-btn @click="currentCmp='about'">About</v-btn>
    <v-btn @click="currentCmp='first'">First Cmp</v-btn>
    <v-btn @click="currentCmp='second'">Second Cmp</v-btn>
    <v-btn @click="quit">Quit</v-btn>

    <about v-if="currentCmp=='about'"></about>
    <first v-if="currentCmp=='first'"></first>
    <second v-if="currentCmp=='second'"></second>
</div>
</template>

<script>
module.exports = {
    data: function() {
        return {
            currentCmp : "about",
            snackbar : {
                show:true,
                text:""
            }
        }
    },
    mounted: function() {
    },
    components: {
        'about': httpVueLoader('/cmp/about/about.vue'),
        'first': httpVueLoader('/cmp/first/first.vue'),
        'second': httpVueLoader('/cmp/second/second.vue')
    },
    methods: {
        quit: function(){
            app.api(
                {method:"GET",url:"/api/action/quit"}
            ).then(data => {
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
