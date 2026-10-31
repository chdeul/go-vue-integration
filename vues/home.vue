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
                    message ( rep.data.Data )

                }).catch( (err) => {
                    this.erreur = "Get /api/exemple: "+ err
                })
            },

           callGetParam() {

                // Example params in Rest route path : /api/exemple/(mot)/print/(nb):
                let mot = "IamFine"
                let nb = 99
                
               // Appel serveur 
                myAxios.get("/exemple/"+mot+"/print/"+nb).then( (rep) => {

                    // Tooltip top right
                    message ( rep.data.Data )

                }).catch( (err) => {
                    this.erreur = "Get /api/exemple/mot/print/#: "+ err
                })
            }, 

           callGetError() {

                // **** Error mot is not matching (\w+) letters only ! ****
                let mot = " I am Fine !"
                let nb = 77
                
               // Appel serveur 
                myAxios.get("/exemple/"+mot+"/print/"+nb).then( (rep) => {

                    // Tooltip top right : not executed
                    message ( rep.data.Data )

                }).catch( (err) => {

                    // Executed on http 404 , because variable mot have whites and !
                    // Dont match the regexp at Go Rest routing 
                    this.erreur = "Get /api/exemple/ I am Fine !/print/77: "+ err
                })
            }, 

 
            callPost() {
            
                // Capitalize field name to allow Go unmarshall
                let payload = { Data: "Hello Go!" }

                // See confirme() into my_global.js tools 

                // Modale dialog
                confirme( "Confirm POST : " , () => { 

                    // Appel serveur 
                    myAxios.post("/exemple", payload).then( (rep) => {

                        message( "envoyé OK" )
                        return true // to hide modal

                    }).catch( (err) => {
                        this.erreur = "Post /api/exemple: "+ err
                    })

                })

            }        
        }

    })    


//# sourceURL=http://localhost:8080/vues/home.vue
</script>

<template id="ui-home">
    <div class="col">
        <h3>Example of golang web application, with simple integration of vue.js</h3>
        <p>
        <b>without</b> needs of nodejs, vite.js, etc.
        </p>
        <br/>
        <h5>REST Go Api tests</h5>
        <div class="d-flex gap-2">
            <button class="btn btn-primary" @click="callGet">Action GET</button>
            <button class="btn btn-success" @click="callGetParam">Action GET Fine</button>
             <button class="btn btn-warning" @click="callPost">Action POST</button>
             <button class="btn btn-danger" @click="callGetError">Action GET error</button>
             <span >{{ erreur}}</span>
        </div>
        <br/>
        <p>Try Api REST endpoints/routes&nbsp;<b> /api/exemple &nbsp; /api/exemple/<i>(ident)</i>/print/<i>(number)</i> </b>&nbsp; nothing else
        </p>
        <br>
        <h5>Scrollup button test</h5>
        <ul>
            <li v-for="n in 100" :key="n">{{ n }} %</li>
        </ul>
        <p>End page  --> see icon to scroll up on rigth</p>
    </div>
</template>

<style>

</style>