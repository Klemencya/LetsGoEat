<script lang="ts">
    import { form, field } from 'svelte-forms';
    import { required } from 'svelte-forms/validators';

    const login = field('login', '', [required()]);
    const password = field('password', '', [required()]);
    const myForm = form(login, password);

</script>

<section class="form">
    <p><b>Your login:</b> <input type="text" bind:value={$login.value} /></p>
    <p><b>Your password:</b> <input type="text" bind:value={$password.value} /></p>

    <button disabled={!$myForm.valid}>Log in</button>
</section>

<style>
    section {
        padding-top: 100px;
        display: flex;
        flex-direction: column;
    }

    .form {
        text-align: left;
        margin-right: 30%;
        margin-left: 30%;
        margin-top: 10%;
        border: 2px solid #9e4eca;
        padding: 50px
    }

    p {
        width: 100%;
    }

    input {
        margin-top: 20px;
        padding: 10px;
        width: 100%;
    }
    button {
        margin-top: 10px;
        width: 50%;
        padding: 10px
    }
</style>
