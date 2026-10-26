<script>

    // Exemple de Mixin :  mixins : [ commons ],    this.formatDate( d ) 
    var commons = {
        methods : {
            formatDate : function(date) {
                if ( date === undefined ) return "?"
                return moment(date).format('DD/MM/YYYY')
            }
        }
    }

//# sourceURL=http://localhost:8080/vues/mixins.vue
</script>