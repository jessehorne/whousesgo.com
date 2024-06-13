var count = document.getElementById("count");
var options = {
    valueNames: ['title', 'location']
};

var l = new List('company-list', options);

l.on('updated', function() {
    count.innerHTML = "(" + l.visibleItems.length + " results)";
});

var companyList = document.getElementById("list");

var newCompanyTemplate = `
<li class="list-item">
  <div class="details">
    <div class="title"><a href="{WEBSITE}">{TITLE} ></a></div>
    <div class="location">{LOCATION}</div>
  </div>
</li>
`

async function updateList() {
    l.clear();

    const res = await fetch("/api/companies");
    const data = await res.json();

    data.companies.forEach(function(c) {
        l.add({
            "title": "<a href='" + c.Website + "'>" + c.Name + " ></a>",
            "location": c.Location
        });
    });

    count.innerHTML = "(" + data.companies.length + " results)";
}

updateList();