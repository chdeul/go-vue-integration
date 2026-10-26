<script>

    const routes = [

      // Accueil
      { path: '/vue/home', component: app.component("ui-home") },

      // Etc.
      { path: '/vue/param', component: app.component("ui-param") },

      // Default : home
      { path: '/', redirect: '/vue/home' },
      { path: '/index.html', redirect: '/vue/home' },
      { path: '/:pathMatch(.*)*', redirect: (to) => `/vue/home?NOT_FOUND=${to.path}` }
     ]

    // Vue router en mode web path  ( not #mark ) 
    var router = VueRouter.createRouter({
      history:  VueRouter.createWebHistory( ContextPath ),
      routes: routes
    })

    // Bind router
    app.use(router)

    // Mode debug : *** TODO change for Production ***
    app.config.devtools = true

    router.isReady().then(() => {

      // Start vuejs app
      app.mount('#app')

     // Sur refresh navigateur, ou favoris ou openid connect : toute l'app recréée 
      const params = new URLSearchParams(window.location.search);
      if ( params && params.size > 0 ) {
          // Chercher query string  ? route=/vue/xxx & yyy=zzz...
          const value = params.get("route");
          if (value && value.startsWith("/vue/") ) {
            // /vue route trouvée , supprime des query
            params.delete( "route")
            // reconstitue les autres param yyy=zzz... eventuels
            qs = params.toString()
            if ( qs ) qs = "?" + qs
            // Route vers cela avec vue router
            router.push( value  + qs )
          }
      } 
 
    })   

//# sourceURL=http://localhost:8080/vues/routes.vue
</script>
