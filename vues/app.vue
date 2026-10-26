<script>

 // DO NOT Touch : Value injected at backend !
 const ContextPath = ''

// --------------------------

    // Config axios, appels ajax (xhr)
    var myAxios = axios.create({ baseURL: ContextPath+'/api'})

    // Transmet le token d'authentication sur chaque requete axios
    myAxios.defaults.withCredentials = true;

    // Traitement standard d'erreur appel ajax axios
    myAxios.interceptors.response.use(
        function(response) {
            return response
        },
        function(error) {
            // Non autorisé
            if (error.response && error.response.status === 401) {
                if ( router.currentRoute.value.path !== "/vue/login" ) {
                   // error treatment : ex redirect to login
                }
            }
            // console.log( "axios intercept: ", error)
            if (error.response && error.response.status >= 500) {
                // error treatment, ex custom error page
            }
            return Promise.reject(error);
        }
    )

// --------------------------

    // Bouton de scrolling arrière si longue page scrollée en bas
    const btnbtt = document.getElementById("backToTop")
    if (btnbtt) {
        window.addEventListener("scroll", () => {
        if (window.scrollY > 120) {
            btnbtt.classList.add("show");
        } else {
            btnbtt.classList.remove("show");
        }
        });
        // remonte 
        btnbtt.addEventListener("click", () => {
            window.scrollTo({ top: 0,behavior: "smooth" });
        });
    }
  
// --------------------------

    // Create vuejs
    var app = Vue.createApp({
        template : "#ui-app",

        // App global data
        // access by this.$root.global_var
        data : function() {
            return {
                global_var : "xxxx",
            }
        },

        created : function () {

        },

        methods : {

        }
    })

// --------------------------

//# sourceURL=http://localhost:8080/vues/app.vue
</script>


<!-- template de la single page -->
<template id="ui-app">
    <div>
        <!-- header + main content -->
        <ui-header></ui-header>
        <div class="container-fluid">
            <div class="row app_padding" >
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<!-- ----------------------- -->

<style>
.app_padding {
    padding-top: 10px; 
    padding-bottom: 10px;
}
</style>

<!-- ----------------------- -->
