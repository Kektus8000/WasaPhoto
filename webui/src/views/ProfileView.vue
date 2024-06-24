<script>
export default{
  data(){
    return{
      errormsg: "",
      profilo: localStorage.getItem('profiloCercato'),
      username: localStorage.getItem('username'),
      identifier: localStorage.getItem('identifier'),
    }
  },
  methods: {
    async refresh(){
      console.log(profilo);
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
      <h4 @click = "() => {this.$router.replace('/session');}" >Torna alla HomePage</h4>
      <h1 style = "font-weight: bold;">Profilo di {{this.profilo.username}}</h1>
      <div class=statistiche>
        <h4>Followers: {{this.profilo.seguaci ? this.profilo.seguaci.length : 0}}</h4>
        <h4>Seguiti: {{this.profilo.seguiti ? this.profilo.seguiti.length : 0}}</h4>
      </div>
    </header>

    <div class = riga>
      <template class = listaFollowers v-for = "utente in this.profilo.fotoPubblicate">
        <div class = colonna>
          <img class = foto :src = utente.link alt="foto">
        </div>
      </template>
    </div>
  </body>
</template>

<style>
  body{
    margin: 0;
    margin-right: 0;
    width:100%;

    background-color: whitesmoke;
  }

  .intestazione{
    width:100%;
    display: flex;
    align-items: center;
    justify-content:space-between;
    cursor: pointer;
    height:100px;

    background-color:brown;
    text-align: left;

    border-bottom-left-radius: 30px;
    border-bottom-right-radius: 30px;
    border-bottom: 3px solid black;
  }

  .riga{
    display: flex;
    flex-wrap: wrap;
  }

  .colonna{
    flex: 1 0 30%;
    padding: 5px;
  }

  .foto{
    width: 100%;
    height:100%;
    border-radius: 10px;;
  }
</style>