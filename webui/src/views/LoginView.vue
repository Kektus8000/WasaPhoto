<template>
  <title>Schermata Login</title>

  <header>
    <h2>Effettua il Login</h2>
  </header>
  
  <hr>

  <main>
    <body class="principale">
      <input placeholder="username" v-model="username">
      <button class = "conferma" 
      :disabled= "username.length > 16 || username.length < 3"
      @click = "trovaUtente"> Effettua Login </button>
    </body>
    

    <div class="fondo">
      {{this.username}}
    </div>
  </main>


</template>

<script>
export default{
  data(){
    return{
      username: "",
      errormsg: ""
    }
  },
  methods:{
    async trovaUtente(){
        try{
            let response = await this.$axios.post("/session", this.username);
            localStorage.setItem('identifier', response);
            this.$router.replace('/userProfile/:userID')
            alert("Success");
        }
        catch(e)
        {
            this.errormsg = e.toString();
            alert("Failure");
        }
    }
  }
}
</script>

<style>
  header{
    text-align: center;
    top:0;
  }

  .principale{
    display: grid;
    align-items:center;
    justify-content: center;
    width: 100%;
  }
  .fondo{
    bottom:0;
    border-top: 1px solid black;
  }

</style>