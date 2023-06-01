<script lang="ts">

</script>

<svelte:head>
  <title>URL Shortner</title>
</svelte:head>

<main>
  <header>
    <div>
      <h1>URL Shortener</h1>
    </div>
  </header>
  <section id="form-section">
    <form action="">
      <input type="text" name="url" id="url" placeholder="Enter URL">
      <input type="submit" id="submit" value="Shorten">
    </form>
  </section>
</main>

<style>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Chakra Petch', sans-serif;
  }

  h1 {
    font-size: 45px;
  }

  #url {
    width: 30%;
    height: 30px;
    font-size: 20px;
  }

  #submit {
    border: none;
    border-radius: 5px;
    background-color: rgb(49, 148, 241);
    color: white;
    padding: 10px;
    font-size: 20px;
  }

  #submit:hover {
    cursor: pointer;
    background-color: rgb(49, 148, 241, 0.8);
  }

  #submit:active {
    transform: scale(0.95);
  }
</style>
