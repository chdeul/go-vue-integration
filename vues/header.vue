<script>
    app.component("ui-header", {
        template: "#ui-header"
    });
//# sourceURL=http://localhost:8080/vues/header.vue
</script>

<template id="ui-header">

  <!-- header sombre -->
  <header class="navbar navbar-expand-sm bg-dark navbar-dark d-print-none">
    <div class="container-fluid">

      <!-- logo , home page -->
      <a class="navbar-brand" href="#"><img src="../resources/img/favicon.png" height="40"/>&nbsp; Exemple </a>

      <!-- compact menu sur petit ecran -->
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#collapsibleNavbar">
        <span class="navbar-toggler-icon"></span>
      </button>

      <!-- menu -->
      <div class="collapse navbar-collapse" id="collapsibleNavbar">

        <!-- menu authentifié , partie à gauche --> 
        <ul class="navbar-nav">

           <!-- menu public, à droite -->
           <li class="nav-item dropdown" >
              <span class="nav-link dropdown-toggle" role="button" data-bs-toggle="dropdown">Menu</span>
              <ul class="dropdown-menu  dropdown-menu-end">
                <li><router-link to="/vue/home" class="dropdown-item">Accueil</router-link></li>
              </ul>
           </li>

        </ul>

        <!-- menu public , partie à droite -->
        <ul class="navbar-nav ms-auto">

           <!-- Authentification -->
           <li class="nav-item dropdown">
              
              <span class="nav-link dropdown-toggle" role="button" data-bs-toggle="dropdown"><i class="fa fa-user"></i> xxxxx </span>
              <ul class="dropdown-menu  dropdown-menu-end">
                <li><router-link to="/vue/param" class="dropdown-item"><i class="fa fa-cog"></i> Paramètres</router-link></li>
                <li><router-link to="/vue/logout" class="dropdown-item"><i class="fa fa-power-off"></i> Déconnexion</router-link></li>
              </ul>
           </li>

        </ul>

      </div>

    </div>
  </header>
</template> 

<style>
.navbar-brand {
  padding-top : 0px;
  padding-bottom : 0px;
}
</style>