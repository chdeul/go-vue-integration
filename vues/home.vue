<script>
    app.component('ui-home', {
        template : "#ui-home",
        mixins : [ commons ],
        data () {
            return {
                data : "the data",
                erreur : ""
            }
        },

        mounted () {
            // Exemple d'appel de mixins
            console.log(  this.formatDate( new Date() )  )
        },
        methods : {
            callGet() {

               // Appel serveur 
                myAxios.get("/exemple").then( (rep) => {

                    // See message() into my_global.js tools 

                    // Tooltip top right
                    message ( "message function : reçu " + rep.data.Data )

                }).catch( (err) => {
                    this.erreur = "Get /api/exemple: error ! "
                })
            },
            callPost() {
            
                // Capitalize field name to allow Go unmarshall
                let payload = { Data: "Hello Go!" }

                // Modale dialog
                confirme( "confirme envoi : " , () => { 

                    // Appel serveur 
                    myAxios.post("/exemple", payload).then( (rep) => {

                        // See confirme() into my_global.js tools 

                        message( "envoyé OK" )
                        return true // to hide modal

                    }).catch( (err) => {
                        this.erreur = "Post /api/exemple: error ! "
                    })

                })

            }        
        }

    })    


//# sourceURL=http://localhost:8080/vues/home.vue
</script>

<template id="ui-home">
    <div class="col">
        <h5>Accueil</h5>
        <p>
        Exemple d'application Golang / Vuejs, sans utilisation de nodejs, vite.js, etc.
        </p>
        <div class="d-flex gap-2">
            <button class="btn btn-primary" @click="callGet">Action GET</button>
            <button class="btn btn-primary" @click="callPost">Action POST</button>
            <span >{{ erreur}}</span>
        </div>
        <br>
        <ul>
            <li v-for="n in 100" :key="n">{{ n }} ....</li>
        </ul>
        <p>End page</p>
    </div>
</template>

<style>

</style>