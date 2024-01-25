<template>
    <body>
      <h1> Effettua il Login</h1>

      <input type = "text" placeholder="Inserisci il tuo username" v-model = "username">

      <button :disabled="username.length < 6 || username.length > 20"
      @click = "cercaUtente"> Login </button>

      <h3 v-show= "username.length < 6 || username.length > 20"> L'username non è della lunghezza adatta </h3>
    </body>
</template>
  
<script>
  export default{
    data() {
      return {
        username : null,
        errormsg : null,
      }
    },
    methods: {
      cercaUtente(){
        try{
            let response = this.$axios.post("/session", this.username);
            router.push()
        }
        catch(e)
        {
            this.errormsg = e.toString();
            console.log(this.errormsg)
        }
        
      }
    }
  }
  </script>
  
  <style>
    
    body {
      justify-content: center;
      height: 200%;
      display: grid;
      place-items: center;
    }
  </style>