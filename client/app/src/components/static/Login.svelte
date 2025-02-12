<script lang="ts">
    let username:string = "";
    let password: string = "";
    console.log(username, password)
</script>



<form action="" >
    <input type="text" name="username" bind:value={username}>
    <input type="password" name="password" bind:value={password}>
    <button type="submit">login</button>
</form>