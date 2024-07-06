<script>

const READER = new FileReader();
const STATOBAN = ["BAN", "UNBAN"];
const STATOFOLLOW = ["FOLLOW", "UNFOLLOW"];

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
      console.log(this.visitor.visitorID);
      console.log(this.profilo.ID);
    },
    ///////////////////////////// FUNZIONI VISITOR ////////////////////////////////////////////
    async visitorInfo(){
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
    async salvaStato(){
      try
      {
        //Aggiornamento del Follow
        if (this.visitor.isFollowing !== this.visitor.initialFollowing)
        {
          // Se il visitatore segue il propritario ma prima no...
          if (this.visitor.isFollowing)
          {
            await this.$axios.post('/userProfile/' + this.visitor.visitorID + '/following/' + this.profilo.ID, {}, {headers: {Authorization: "Bearer " + this.visitor.visitorID}} );
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
            await this.$axios.post('/userProfile/' + this.visitor.visitorID  + '/banList/' + this.profilo.ID, {}, {headers: {Authorization: "Bearer " + this.visitor.visitorID }} );

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
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
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
            foto.File = URL.createObjectURL(temp.data);
          }  
        }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async uploadPhoto(){
      try
      {
        let foto = document.getElementById('upload');
        READER.readAsArrayBuffer(foto.files[0]);
        READER.onload = async () => {
          let response = await this.$axios.put('/userProfile/${this.visitorID}/publishedPhotos/', READER.result, { headers: {Authorization: "Bearer " + this.visitor.visitorID} });
        };
        alert("Foto pubblicata!");
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async deletePhoto(photoID){
      try
      {
        await this.$axios.delete('/userProfile/' + this.profilo.ID + '/publishedPhotos/' + photoID,
        { headers: {Authorization: "Bearer " + this.profilo.ID} });
        alert("Foto cancellata");
        this.refresh();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async mostraDialog(){ document.getElementById("cambiaNome").style.display = "inline"; },
    
    async chiudiDialog(){
      this.newUsername = "";
      document.getElementById("cambiaNome").style.display = "none"; 
    },
    
    async cambiaUsername(){
      try
      {
        let response = await this.$axios.post('/user/' + this.profilo.ID + '/username', {username: this.newUsername}, { headers: {Authorization: "Bearer " + this.visitor.visitorID} })
        this.profilo.nome = this.newUsername;
        localStorage.setItem('username', this.newUsername);
        this.newUsername = "";
        alert("Username cambiato!");
        this.chiudiDialog();
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async controllaBannati(){ this.$router.push({path: '/userProfile/'+ this.profilo.ID + '/banList'});},
    async tornaHomePage(){
      if (this.visitor.visitorID != this.profilo.ID) {this.salvaStato();}
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
        <button @click = "cambiaUsername" :disabled= "this.newUsername.length > 16 || this.newUsername.length < 3">Cambia</button>
        <button @click = "chiudiDialog" > Chiudi </button>
      </div>
    </dialog>

    <div class = fotoCondivise>
      <div v-for = "immagine in this.profilo.fotoPubblicate">
        <img class = foto-profilo :src = "immagine.File" :key = immagine.PhotoID>
        <button class = photo-delete-button 
        @click= deletePhoto(immagine.PhotoID)
        v-if = "Number(this.profilo.ID) === Number(this.visitor.visitorID)">Cancella </button>
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
    position: relative;
    transform: translate(+390%, -120%);
    border-radius: 10px;
    border: 2px solid white;
    height:40px;
    width: 100px;

    background-color: rgb(178, 34, 34);
    color:white;
  }
</style>