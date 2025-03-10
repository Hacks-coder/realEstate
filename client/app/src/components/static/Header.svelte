<script>
	import Text from "../resuable/Text.svelte";
    import IconMenu from '~icons/lucide/menu'
    let isNavigationVisible = false;
    const toogleNavigation = () => {
        isNavigationVisible = !isNavigationVisible;
    }
    const navigation = [
        {href: "/home", child: "Home"},
        {href: "/home", child: "About"},
        {href: "/home", child: "Contact"},
        {href: "/home", child: "Profile"},
        {href: "/homem", child: "Buy"},
    ]
</script>
<header class="flex justify-between  p-5">
    <div>
        <Text text =  "hacks_coder" className = ""/>
    </div>
        <nav class="md:hidden">
            <IconMenu onclick={toogleNavigation} class="text-2xl"/>
            {#if isNavigationVisible}
                {#each navigation as item}
                    <div class="flex gap-2">
                        <a href={item.href} class="text-lg mt-2">{item.child}</a>
                    </div>
                {/each}
            {/if}
        </nav>
    <nav class="hidden md:flex gap-5">
       {#each navigation as item}
            <div class="flex gap-2 text-xl">
                <a href={item.href}>{item.child}</a>
            </div>
        {/each}
    </nav>
</header>