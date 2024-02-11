catch(e)
<template>
  <title>Schermata Login</title>

  <header>
    <h2>Effettua il Login</h2>
  </header>
  
  <hr>

  <main>
    <body class="principale">
      <input type ="text" placeholder="Inserisci il tuo username" v-model="username">
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
        console.log(this.username);
        try{
            let response = await this.$axios.post("/session",{username: this.username});
            var profilo = response.data;
            localStorage.setItem('identifier', profilo.UserID);
            localStorage.setItem('username', profilo.Username);
            this.$router.replace('/userProfile/:userID');
        }
        catch(e)
        {
            this.errormsg = e.toString();
            alert(this.errormsg);
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