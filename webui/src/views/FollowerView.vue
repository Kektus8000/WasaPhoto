<script>
export default{
  data(){
    return{
      visitorID : Number(localStorage.getItem('identifier')),
      ID : localStorage.getItem('IDCercato'),
      seguaci : JSON.parse(localStorage.getItem('SeguaciSessione')),
      bannati : JSON.parse(localStorage.getItem('BannatiSessione'))
    }
  },
  methods:{
    async refresh(){},
    async banUser(banned){
      try
      {
        let indice = this.seguaci.map(user => user.UserID).indexOf(banned.UserID);
        this.seguaci.splice(indice, 1);
        alert("Hai bloccato " + banned.Username);
        localStorage.setItem('SeguaciSessione', JSON.stringify(this.seguaci));

        await this.$axios.put('/userProfile/' + this.ID + '/banList/' + banned.UserID, {}, {headers: {Authorization: "Bearer " + this.visitorID}} );

        var temp = {UserID : banned.UserID, Username: banned.Username};
        if (this.bannati == null) {this.bannati = temp;}
        else {this.bannati.push(temp);}
        localStorage.setItem('BannatiSessione', JSON.stringify(this.bannati));
        this.refresh();
      }
      catch(e)
      {
        if (e.response != null && e.response.status == 500) {alert("Un errore nel server impedisce l'operazione!");}
      }
    },
    async visitaProfilo(checkID){
      localStorage.setItem('IDCercato', checkID);
      this.$router.push({path: '/userProfile/' + checkID });
    },
  },
  mounted (){
    this.refresh();
  }
}
</script>

<template>
  <body>
    <header class=intestazione-seguaci>
      <h2 style = "font-weight: bold; padding-left: 20px; cursor:pointer;"
      @click = "() => {this.$router.back();}"> Torna indietro </h2>
      <h1 style = "font-weight: bold;">Account Seguaci</h1>
      <h2 style = "padding-right: 20px;"> Seguaci : {{this.seguaci != null ? this.seguaci.length : 0}}</h2>
    </header>

    <section class = lista-seguaci>
      <div class = follower v-for = "utente in this.seguaci" :key = "utente.UserID">
        <h2 style = "cursor: pointer; font-weight: bold;" @click = visitaProfilo(utente.UserID)>{{utente.Username}}</h2>
        <button class = blocca style = "color:red" @click = banUser(utente) show = "this.ID === Number(localStorage.getItem('identifier'))">Ban</button>
      </div>
    </section>
  </body>
</template>

<style>
  body{
    margin: 0;
    margin-right: 0;
    width:100%;

    display: inline-block;
    background-color: whitesmoke;
  }

  .intestazione-seguaci{
    position: fixed;
    width:100%;
    height:100px;

    display: flex;
    align-items:center;
    justify-content: space-between;

    background-color: cadetblue;
    border-bottom: 5px solid black;
  }

  .lista-seguaci{
    padding-left: 10%;
    padding-right: 10%;
    padding-top:120px;
    width: 100%;
    
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }

  .follower{
    cursor: pointer;
    border-radius: 10px;
    box-shadow: 10px 10px 5px cadetblue;
    text-align: center;
    border: 2px solid black;
  }

</style>