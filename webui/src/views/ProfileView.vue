<script>

const READER = new FileReader();
const STATOBAN = ["BAN", "UNBAN"];
const STATOFOLLOW = ["FOLLOW", "UNFOLLOW"];
const STATOLIKE = ["Mi Piace", "Ti Piace"];

export default{
  data(){
    return{
      newUsername : "",

      errormsg: "",
      
      visitor : {
        visitorID : Number(localStorage.getItem('identifier')),
        visitorNome : localStorage.getItem('username'),
        visitorSeguaci : JSON.parse(localStorage.getItem('SeguaciSessione')),
        visitorSeguiti : JSON.parse(localStorage.getItem('SeguitiSessione')),
        visitorBannati : JSON.parse(localStorage.getItem('BannatiSessione')),
        initialFollowing : false,
        initalBanned : false,
        hasBanned : false,
        isFollowing : false
      },

      profilo : {
        ID : Number(localStorage.getItem('IDCercato')),
        nome : "",
        seguaci : [],
        seguiti : [],
        bannati : [],
        fotoPubblicate : []
      }
    }
  },
  methods: {
    async refresh(){
      await this.recuperaInfo();
      if (this.visitor.visitorID != this.profilo.ID) {this.visitorInfo();}
    },
    ///////////////////////////// FUNZIONI VISITOR ////////////////////////////////////////////
    visitorInfo(){
      if (this.visitor.visitorSeguiti != null) {
        this.visitor.isFollowing = this.visitor.visitorSeguiti.some(e => Number(e.UserID) === this.profilo.ID);
        this.visitor.initialFollowing = this.visitor.isFollowing;
      }
      if (this.visitor.visitorBannati != null) {
        this.visitor.hasBanned = this.visitor.visitorBannati.some(e => Number(e.UserID) === this.profilo.ID);
        this.visitor.initalBanned = this.visitor.hasBanned;
      }

      document.getElementById('followButton').value = STATOFOLLOW[0];
      document.getElementById('banButton').value = STATOBAN[0];
      if (this.visitor.isFollowing) {document.getElementById('followButton').value = STATOFOLLOW[1];}
      if (this.visitor.hasBanned) {document.getElementById('banButton').value = STATOBAN[1];}
    },
    gestioneFollow(){
      this.visitor.isFollowing = !this.visitor.isFollowing;
      if (this.visitor.isFollowing) {document.getElementById('followButton').value = STATOFOLLOW[1];}
      else {document.getElementById('followButton').value = STATOFOLLOW[0];}
      
      if (this.visitor.hasBanned) {
        this.visitor.hasBanned = false;
        document.getElementById('banButton').value = STATOBAN[0];
      }
    },
    gestioneBan(){
      this.visitor.hasBanned = !this.visitor.hasBanned;
      if (this.visitor.hasBanned) {document.getElementById('banButton').value = STATOBAN[1];}
      else {document.getElementById('banButton').value = STATOBAN[0];}
      
      if (this.visitor.isFollowing) 
      {
        this.visitor.isFollowing = false;
        document.getElementById('followButton').value = STATOFOLLOW[0];
      }
    },
    async salvaLike()
    {
      try
      {
        for (let i = 0; i < this.profilo.fotoPubblicate.length; i++)
        {
          var foto = this.profilo.fotoPubblicate[i];
          if (foto.isLiked !== foto.initiallyLiked)
          {
            if (foto.isLiked) { 
              await this.$axios.put('/userProfile/' + this.visitor.visitorID + '/stream/' + foto.PhotoID + '/likes/', {}, { headers: {Authorization: "Bearer " + this.visitor.visitorID} });
              //Questo serve per PhotoView principalmente
              var newLike = {
                UserID : this.visitor.visitorID,
                Username : this.visitor.visitorNome
              };
              if (foto.Likes != null) {foto.Likes.push(newLike);}
              else {foto.Likes = [newLike];}
            }
            else {
              await this.$axios.delete('/userProfile/' + this.visitor.visitorID + '/stream/' + foto.PhotoID + '/likes/', { headers: {Authorization: "Bearer " + this.visitor.visitorID} });
              //Questo serve per PhotoView principalmente
              let indice = foto.Likes.map(user => user.UserID).indexOf(this.visitor.visitorID);
              foto.Likes.splice(indice, 1);
            }
          }
        }
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 403:
              alert("Non hai i permessi per completare l'operazione!");
              break;
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        } 
      }
    },
    async salvaStato(){
      try
      {
        //Aggiornamento del Follow
        if (Number(this.visitor.visitorID) != Number(this.profilo.ID))
        {
          if (this.visitor.isFollowing !== this.visitor.initialFollowing)
          {
            // Se il visitatore segue il propritario ma prima no...
            if (this.visitor.isFollowing)
            {
              await this.$axios.put('/userProfile/' + this.visitor.visitorID + '/following/' + this.profilo.ID, {}, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} );
              var temp = {UserID : this.profilo.ID, Username: this.profilo.nome};
              if (this.visitor.visitorSeguiti == null) {this.visitor.visitorSeguiti = [temp];}
              else {this.visitor.visitorSeguiti.push(temp);}
            }
            // Se il visitatore seguiva il proprietario e ora non più...
            else
            {
              await this.$axios.delete('/userProfile/' + this.visitor.visitorID + '/following/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} ); 
              let indice = this.visitor.visitorSeguiti.map(user => user.UserID).indexOf(this.profilo.ID);
              this.visitor.visitorSeguiti.splice(indice, 1);
            }
          }
          //Aggiornamento del Ban
          if (this.visitor.hasBanned !== this.visitor.initalBanned)
          {
            // Se il visitatore ha bannato il proprietario...
            if (this.visitor.hasBanned)
            {
              await this.$axios.put('/userProfile/' + this.visitor.visitorID  + '/banList/' + this.profilo.ID, {}, {headers: {Authorization: "Bearer " + this.visitor.visitorID }} );

              // L'utente bannato viene rimosso dai seguiti (se vi è)...
              if (this.visitor.visitorSeguiti != null && this.visitor.visitorSeguiti.some(fl => Number(fl.UserID) === this.profilo.ID))
              {
                let indice = this.visitor.visitorSeguiti.map(user => user.UserID).indexOf(this.profilo.ID);
                if (indice != -1) {this.visitor.visitorSeguiti.splice(indice, 1);}
              }
              // ...e viene aggiunto ai bannati
              var temp = {UserID : this.profilo.ID, Username: this.profilo.nome};
              if (this.visitor.visitorBannati == null) {this.visitor.visitorBannati = [temp];}
              else {this.visitor.visitorBannati.push(temp);}
            }
            else
            {
              await this.$axios.delete('/userProfile/' + this.visitor.visitorID + '/banList/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} );
            
              // L'utente viene rimosso dai bannati
              let indice = this.visitor.visitorBannati.map(user => user.UserID).indexOf(this.profilo.ID);
              this.visitor.visitorBannati.splice(indice, 1);
            }
          }
          // Aggiornamento del localStorage
          localStorage.setItem('SeguitiSessione', JSON.stringify(this.visitor.visitorSeguiti));
          localStorage.setItem('BannatiSessione', JSON.stringify(this.visitor.visitorBannati));
        }
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 403:
              alert("Non hai i permessi per completare l'operazione!");
              break;
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        } 
      }
    },
    fotoPiaciuta(photoID)
    {
      var foto = this.profilo.fotoPubblicate.find(foto => foto.PhotoID === photoID);
      return foto.isLiked ? STATOLIKE[1] : STATOLIKE[0];
    },
    likePhoto(photoID)
    {
      var foto = this.profilo.fotoPubblicate.find(foto => foto.PhotoID === photoID);
      foto.isLiked = !foto.isLiked;
      if (foto.isLiked) {foto.likeCount += 1}
      else {foto.likeCount -= 1}
    },
    async infoFoto(photoID)
    {
      await this.salvaLike();
      await this.salvaStato();
      var foto = this.profilo.fotoPubblicate.find(foto => Number(foto.PhotoID) === Number(photoID));
      localStorage.setItem('FotoAnalizzata', JSON.stringify(foto));
      this.$router.push({path: '/userProfile/' + this.profilo.ID + '/publishedPhotos/' + photoID});
    },
    ////////////////////////////// FUNZIONI OWNER //////////////////////////////////////////////
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.profilo.ID, {headers: {Authorization: "Bearer " + this.visitor.visitorID}});
        
        this.profilo.nome = response.data.Username;

        if (response.data.Followers != null) {this.profilo.seguaci = response.data.Followers; }
        if (response.data.Followings != null) {this.profilo.seguiti = response.data.Followings; }
        if (response.data.Banneds != null) {this.profilo.bannati = response.data.Banneds;}
        if (response.data.PublishedPhotos != null)
        {
          this.profilo.fotoPubblicate = response.data.PublishedPhotos;
          for (let i = 0; i < this.profilo.fotoPubblicate.length; i++)
          {
            var foto = this.profilo.fotoPubblicate[i];
            let temp = await this.$axios.get('/userProfile/' + this.profilo.ID + '/publishedPhotos/' + foto.PhotoID, 
            {responseType: 'blob',
            headers: {Authorization: "Bearer " + this.visitor.visitorID}});
            foto.Path = URL.createObjectURL(temp.data);

            foto.isLiked = false;
            if (foto.Likes != null) {foto.isLiked = foto.Likes.some(like => like.UserID === this.visitor.visitorID);}
            foto.initiallyLiked = foto.isLiked;
            foto.likeCount = foto.Likes != null ? foto.Likes.length : 0;
          }  
        }
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 403:
              alert("Non puoi accedere al profilo di questo utente perché sei stato bannato da lui!");
              break;
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        }
      }
    },
    async uploadPhoto(){
      try
      {
        let img = document.getElementById('upload');
        READER.readAsArrayBuffer(img.files[0]);
        READER.onload = async () => {
          let response = await this.$axios.post('/userProfile/' + this.profilo.ID + '/publishedPhotos/', READER.result, { headers: {Authorization: "Bearer " + this.visitor.visitorID} });
          
          var foto = {
            PhotoID : response.data,
            PublisherID : this.profilo.ID,
            PublisherName : this.profilo.nome,
            File : "/tmp/userProfile/" + this.profilo.ID + "/publishedPhotos/" + response.data,
            Comments : [],
            Likes : [],
            isLiked : false,
            likeCount: 0
          }
          
          let temp = await this.$axios.get('/userProfile/' + this.profilo.ID + '/publishedPhotos/' + foto.PhotoID, 
          {responseType: 'blob',
          headers: {Authorization: "Bearer " + this.visitor.visitorID}});
          foto.Path = URL.createObjectURL(temp.data);

          //La foto viene messa in cima alla lista di Foto pubblicate
          if (this.profilo.fotoPubblicate) {this.profilo.fotoPubblicate.unshift(foto);}
          else {this.profilo.fotoPubblicate = foto;}
        };
        alert("Foto pubblicata!");

      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        }
      }
    },
    async deletePhoto(photoID){
      try
      {
        await this.$axios.delete('/userProfile/' + this.profilo.ID + '/publishedPhotos/' + photoID,
        { headers: {Authorization: "Bearer " + this.profilo.ID} });
        alert("Foto cancellata");
        await this.salvaLike();
        await this.salvaStato();
        this.refresh();
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 403:
              alert("Soltanto il proprietario della foto può cancellarla!");
              break;
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        }
      }
    },
    async mostraDialog(){ document.getElementById("cambiaNome").style.display = "inline"; },
    
    async chiudiDialog(){
      this.newUsername = "";
      document.getElementById("cambiaNome").style.display = "none"; 
    },
    
    async setMyUsername(){
      try
      {
        await this.$axios.put('/user/' + this.profilo.ID + '/username', {username: this.newUsername}, { headers: {Authorization: "Bearer " + this.visitor.visitorID} })
        this.profilo.nome = this.newUsername;
        localStorage.setItem('username', this.newUsername);
        this.newUsername = "";
        alert("Username cambiato!");
        this.chiudiDialog();
      }
      catch(e)
      {
        if (e.response != null)
        {
          switch (e.response.status)
          {
            case 403:
              alert("L'username scelto è già in uso da un altro utente!");
              break;
            case 404:
              alert("Errore nell'autenticazione!");
              break;
            case 500:
              alert("Un errore nel server impedisce l'operazione!");
              break;
          }
        }
      }
    },
    async controllaBannati(){
      await this.salvaLike();
      await this.salvaStato();
      this.$router.push({path: '/userProfile/'+ this.profilo.ID + '/banList'});
    },
    async tornaHomePage(){
      await this.salvaLike();
      await this.salvaStato();
      localStorage.removeItem('IDCercato');
      this.$router.replace('/session');
    }
  },
    mounted(){
      this.refresh();
  }
}
</script>

<template>
  <head>
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
  </head>
  <body>
    <header class=intestazione>
      <h4 @click = "tornaHomePage" style= "padding-left : 10px; font-weight: bold;">Torna alla HomePage</h4>
      <h1 style = "font-weight: bold;">Profilo di {{this.profilo.nome}}</h1>
      <div class=statistiche style= "padding-right : 10px; ">
        <h4>Seguaci: {{this.profilo.seguaci ? this.profilo.seguaci.length : 0}}</h4>
        <h4>Seguiti: {{this.profilo.seguiti ? this.profilo.seguiti.length : 0}}</h4>
      </div>
    </header>

    <dialog id = "cambiaNome" close>
      <h3 style = "text-align: center;"> Scrivi il tuo nuovo username </h3>
      <div style = "align-self: center;">
        <input placeholder="Nuovo username" v-model = this.newUsername style = "text-align: center;">
        <button @click = "setMyUsername()" :disabled= "this.newUsername.length > 16 || this.newUsername.length < 3">Cambia</button>
        <button @click = "chiudiDialog" > Chiudi </button>
      </div>
    </dialog>

    <div class = fotoCondivise>
      <div v-for = "immagine in this.profilo.fotoPubblicate" :key = immagine.PhotoID>
        <img class = foto-profilo :src = "immagine.Path" @click = infoFoto(immagine.PhotoID)>
        <div class = ops-foto style = "display: flex; justify-content: space-between; transform: translate(0%, -110%);">
          <button class = like-profilo  @click = likePhoto(immagine.PhotoID)> <i class="material-icons">favorite</i> {{fotoPiaciuta(immagine.PhotoID)}} : {{immagine.likeCount}} </button>
          <button class = photo-delete-button 
          @click= deletePhoto(immagine.PhotoID)
          v-if = "Number(this.profilo.ID) === Number(this.visitor.visitorID)">Cancella </button>
        </div>
      </div>
    </div>
    
    <div class = footer>
      <div class = ownerActions v-if = "this.profilo.ID == this.visitor.visitorID">
        <h4 style= "padding-left : 100px; font-weight: bold;" @click = "controllaBannati"> Bannati: {{this.profilo.bannati ? this.profilo.bannati.length : 0}}</h4>
        <input id = "upload" type="file" accept="image/*" @change="uploadPhoto" style= "padding-left : 100px;">
        <h4 style= "padding-right : 100px; font-weight: bold;" @click = "mostraDialog()"> Cambia username </h4>
      </div>

      <div class = visitorActions v-else>
        <input  id = banButton style = "font-size: 30px; color:red;" type="button" runat="server" @click="gestioneBan()">
        <input  id = followButton style = "font-size: 30px;" type="button" runat="server" @click="gestioneFollow()">
      </div>
    </div>
  </body>
</template>

<style>

  dialog[open]{
    top:  50%;
    left: 50%;
    display: inline;
    transform: translate(-100%, -60%);
  }

  body{
    margin: 0;
    margin-right: 0;
    width:100%;
    background-color: whitesmoke;
  }

  .intestazione{
    position: fixed;
    width:100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: pointer;
    height:100px;

    background-color: rgb(0, 117, 94);

    border-bottom-left-radius: 30px;
    border-bottom-right-radius: 30px;
    border-bottom: 3px solid black;
  }

  .fotoCondivise{
    padding-top:100px;
    padding-bottom:100px;
    padding-left: 10%;
    padding-right: 10%;
    width: 100%;
    
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }

  .foto-profilo{
    width: 100%;
    height: 300px;
    object-fit: cover;
    border-radius: 15px;
    border: 2px solid black;
  }

  .footer{
    position:fixed;
    width:100%;

    cursor: pointer;
    height:100px;
    bottom: 0;

    background-color:rgb(0, 117, 94);

    border-top-left-radius: 30px;
    border-top-right-radius: 30px;
    border-top: 3px solid black;
  }

  .ownerActions{
    padding-top: 25px;
    display: flex;
    align-items: center;
    justify-content:space-between;
  }

  .visitorActions{
    padding-top: 25px;
    display: flex;
    align-items: center;
    justify-content:space-evenly;
  }

  .photo-delete-button{
    border-radius: 10px;
    border: 2px solid white;
    height:40px;
    width: 100px;

    background-color: rgb(178, 34, 34);
    color:white;
  }

  .like-profilo{
    font-size: 20px;

    border-radius:10px;
    border: 1px solid black;
    color: rgb(80, 200, 120);
    cursor: pointer;
    background-color: white;
  }

  .like-profilo:hover{
    color: white;
    background-color: rgb(80, 200, 120);
  }
</style>