async function fetchIPAddress(domain) {
    const response = await fetch(`https://cloudflare-dns.com/dns-query?name=${domain}&type=A`, {
        headers: {
            'accept': 'application/dns-json'
        }
    });
    const responseObject = await response.json();

    console.log(responseObject);
}


const domain = 'google.com';
const ipAddress = fetchIPAddress(domain);
