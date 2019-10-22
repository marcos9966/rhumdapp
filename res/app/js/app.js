var app = {
    api : function(options) {
        return new Promise((resolve, reject)=>{
            var params={}
            if (options.params) params = options.params
            axios({
                method: options.method,
                url: options.url,
                params: params,
                responseType: 'json'
            }).then(function (response) {
                resolve(response.data);
            }).catch(function (error) {
                if (error.response) {
                    reject(error.response)
                } else if (error.request) {
                    reject(error.request)
                } else {
                    reject(error.message)
                }
            });
        });
    }
};
