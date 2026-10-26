function message(msg, duration = 3000) {

  // Tooltip 
  const box = document.createElement("div")
  box.textContent = msg
  box.classList.add( "message" )
  document.body.appendChild(box)

  // Fade-in
  requestAnimationFrame(() => {
    box.style.opacity = "1"
  })

  // Disparition automatique
  setTimeout(() => {
    box.style.opacity = "0"
    setTimeout(() => box.remove(), 300)
  }, duration)
}

function confirme(msg,callback) {

    // Ajouter le div dans le body
  let overlay = document.getElementById("overlay")
  if ( !overlay ) {
    overlay = document.createElement("div")   
    overlay.id = "overlay"
    document.body.appendChild( overlay )
  }

  // Créer le div
  const modale = document.createElement("div")
  modale.classList.add( "modale" )
 
  // Créer le paragraphe
  const p = document.createElement("p")
  p.textContent = msg

  // Créer les boutons
  const btnYes = document.createElement("button")
  btnYes.textContent = "Oui"
  btnYes.onclick = function () {
     // Disparition
     overlay.style.display = "none"
     modale.remove()
     // Action
     return callback ? callback() : true 
  }

  const btnNo = document.createElement("button")
  btnNo.textContent = "Non"
  btnNo.onclick = function () {
     // Disparition
     overlay.style.display = "none"
     modale.remove()
     return false
  }
  // Ajouter les éléments dans le div
  modale.appendChild(p)
  modale.appendChild(btnYes)
  modale.appendChild(btnNo)

  // Ajouter et afficher
  overlay.appendChild(modale)
  overlay.style.display = "flex" 
}

