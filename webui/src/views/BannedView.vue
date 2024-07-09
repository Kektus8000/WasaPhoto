<script>
export default{
  data(){
    return{
      visitorID : Number(localStorage.getItem('identifier')),
      ID : Number(localStorage.getItem('IDCercato')),
      bannati : JSON.parse(localStorage.getItem('BannatiSessione'))
    }
  },
  methods:{
    async refresh(){
      console.log(this.bannati);
    },
    async unBanUser(banned){
      try
      {
        await this.$axios.delete('/userProfile/' + this.ID + '/banList/' + banned.UserID, {headers: {Authorization: "Bearer " + this.visitorID}} );
        let indice = this.bannati.map(user => user.UserID).indexOf(banned.UserID);
        this.bannati.splice(indice, 1);
        alert(banned.Username + " non è più bannato!");
        localStorage.setItem('BannatiSessione', JSON.stringify(this.bannati));
        this.refresh();
      }
      catch(e)
      {
        if (e.response != null && e.response.status == 500)
        {
          alert("Un errore nel server impedisce l'operazione!");
        }
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
    <header class=intestazione-bannati>
      <h2 style = "font-weight: bold; padding-left: 20px; cursor:pointer;"
      @click = "() => {this.$router.back();}"> Torna indietro </h2>
      <h1 style = "font-weight: bold;">Account Bannati</h1>
      <h2 style = "padding-right: 20px;"> Bannati : {{this.bannati != null ? this.bannati.length : 0}}</h2>
    </header>

    <section class = lista-bannati>
      <div class = utente-bannato v-for = "utente in this.bannati" :key = "utente.UserID">
        <h2 style = "cursor: pointer; font-weight: bold;" @click = visitaProfilo(utente.UserID)>{{utente.Username}}</h2>
        <button class = blocca style = "color:red" @click = unBanUser(utente) show = "this.ID === Number(localStorage.getItem('identifier'))"> Unban </button>
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

  .intestazione-bannati{
    position: fixed;
    width:100%;
    height:100px;

    display: flex;
    align-items:center;
    justify-content: space-between;

    background-color: rgb(150, 0, 24);
    border-bottom: 5px solid black;
  }

  .lista-bannati{
    padding-left: 10%;
    padding-right: 10%;
    padding-top:120px;
    width: 100%;
    
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }

  .utente-bannato{
    cursor: pointer;
    border-radius: 10px;
    box-shadow: 10px 10px 5px rgb(150, 0, 24);
    text-align: center;
    border: 2px solid black;
  }

</style>