<script>
    import Text from "../resuable/Text.svelte";

    import IconRiPriceTag3Line from '~icons/ri/price-tag-3-line';
    import IconFaSolidUsers from '~icons/fa-solid/users';
    import IconMaterialSecurity from '~icons/material-symbols/security';
    import IconFaRegLightbulb from '~icons/fa-regular/lightbulb';
    import IconMdiHandshakeOutline from '~icons/mdi/handshake-outline';
    import IconFaSolidUserTie from '~icons/fa-solid/user-tie';


    const chooseUs = [
        {
            icon: IconRiPriceTag3Line,
            title: "Best Price",
            description: "We offer the most competitive prices, ensuring you get the best value for your investment."
        },
        {
            icon: IconFaSolidUsers,
            title: "Community",
            description: "We are deeply rooted in the local community, providing you with valuable insights and access to the best neighborhoods and opportunities."
        },
        {
            icon: IconMaterialSecurity,
            title: "Secure",
            description: "We prioritize your safety with secure practices, ensuring every transaction is reliable and risk-free."
        },
        {
            icon: IconFaRegLightbulb,
            title: "Guidance",
            description: "We provide professional advice and insights to help you make informed decisions, ensuring a smooth and successful real estate journey."
        },
        {
            icon: IconMdiHandshakeOutline,
            title: "Trusted",
            description: "With a reputation for honesty and reliability, our clients trust us to deliver exceptional results every time."
        },
        {
            icon: IconFaSolidUserTie,
            title: "Agent",
            description: "Our experienced agents offer personalized service, local market expertise, and strong negotiation skills to ensure a seamless and successful real estate experience."
        }
    ]
</script>


    <p class="font-bold text-center mt-10">WHY CHOOSE US?</p>   
<section class="grid grid-cols-[160px_160px] gap-3 grid-rows-[190px_190px] mt-10 md:grid-cols-[300px_300px] lg:grid-cols-[350px_350px_350px] lg:grid-rows-[250px_250px]">
    {#each chooseUs as {title, icon, description}}
        <div class="text-center p-2 flex flex-col items-center  bg-standard rounded-lg shadow-sm shadow-black">
            <div class="w-10 h-10 rounded-full shadow-sm shadow-black flex items-center justify-center"><svelte:component this={icon} class="text-2xl"/></div>
            <Text text = {title} className = "font-bold text-lg mt-3 md:text-xl lg:text-2xl" />
            <Text text = {description} className = "mt-3 text-xxs md:text-md lg:text-lg"/>
        </div>
    {/each}
</section>