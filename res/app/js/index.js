Vue.use(Vuetify, {
    iconfont: 'fa5'
});

new Vue({
    vuetify: new Vuetify(),
    data: {
    },
    mounted: function() {
    },
    components: {
        'home-page': httpVueLoader('/cmp/home/home.vue')
    },
    methods: {
        close: function() {
            myApp.api({
                url:"/api/quit",
            })
            .then(data => console.log("quit") )
            .catch(error => console.log(error));
        }
    }
}).$mount('#app');
