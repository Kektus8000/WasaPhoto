<script>
export default{
  data(){
    return{
      ID : localStorage.getItem('IDCercato'),
      bannati: []
    }
  },
  methods:{
    async refresh(){
      this.recuperaInfo();
    },
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.ID, {headers: {Authorization: "Bearer " + this.ID}});
        
        if (response.data.Banneds != null) {this.bannati = response.data.Banneds; }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
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
      <h2 style = "padding-right: 20px;"> Bannati : {{this.bannati.length}}</h2>
    </header>

    <section class = lista-bannati>
      <div class = utente-bannato v-for = "utente in this.bannati">
        <h2>{{utente}}</h2>
        <button class = blocca style = "color:red"> Unban </button>
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