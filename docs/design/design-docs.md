# Messages

| Name | Description | Synchronous (S) or Asynchronous (A) |
|---|---|---|
| fixturelist.generate | Request a fixture list for the coming season to be generated | A | 
| info.fixturecompleted | A specific fixture has been completed | A | 
| info.playertransferredin | Notification that a new player has been trasferred into a team | A | 
| info.playertransferredout | Notification that a new player has been trasferred out of a team | A | 
| info.seasoncompleted | Notification that indicates the final fixture in a season has been completed | A | 
| info.teamcreated | Notification that a new team has been added | A | 
| info.teamupdated | Notification that a team has been updated | A | 
| info.fixturelistgenerated | Notification that the fixture list for a specific team has been generated | A | 
| leaguetable.updatedata | Update data in the league table data store (team names etc) | S | 
| leaguetable.updateresults | A request to update the league table (normally called after a weeks games have completed) | A | 
| leaguetable.updatestats | A request to update the league stats (normally called after a weeks games have completed) | A | 
| player.add | Add a new player to a team | S | 
| player.delete | Delete a player from a team | S | 
| player.update | Update a players details against a team | S | 
| result.create | Create a new fixture result | A | 
| sponsor.create | Create a new sponsor | S | 
| sponsor.distribute | Begin the distribution of sponsorships | A | 
| sponsor.update | Update an existing sponsor | S | 
| team.create | Create a new team | S | 
| team.relegate | Relegate a team from the league | A | 
| team.update | Update a team in the league | S | 
| transfer.in | Transfer a player in | S | 
| transfer.out | Transfer a player out | S | 


# Message Flows

| Name | Message Flow |
|---|---|
| New Team Added | 1. team.create 2. info.teamcreated |
| Team Updated | 1. team.update 2. info.teamupdated 3. leaguetable.updatedata |
| Fixture list generation | 1. fixturelist.generate 2. info.fixturelistgenerated |
| Season completed | 1. leaguetable.updateresults 2. info.seasoncompleted 3. sponsor.distribute 4. team.relegate |
| Result Added | 1. result.create 2. info.fixturecompleted 3. leaguetable.updateresults 4. leaguetable.updatestats |
| New Sponsor Created | 1. sponsor.create |
| Player transferred in | 1. transfer.in 2. info.playertransferredin |
| Player transferred out | 1. transfer.out 2. info.playertransferredout |
| Sponsor updated | 1. sponsor.update |


# Contexts

| Name | Description | Sends | Receives |
|---|---|---|---|
| front | Handles external HTTP requests and sits behind a load balancer |  |  |
| identity | Handles authentication and identity management |  | info.teamcreated // info.teamupdated // info.teamrelegated // |
| team-manager | Handles all activities around the teams themselves including index data and player management | info.teamcreated // info.teamupdated // transfer.in // transfer.out // | team.create // team.update // player.add // player.update // player.delete // |
| sponsor-manager | Handles all storage, distribution and management of sponsorship deals |  | info.seasoncompleted // sponsor.create // sponsor.update // sponsor.distribute // |
| fixture-manager | Handles scheduling and storage of fixtures and results | info.seasoncompleted // info.resultcompleted // info.fixturelistgenerated // leaguetable.updatedata // team.relegate // | fixturelist.generate // result.create // leaguetable.updateresults // |
| transfer-manager | Manages transfers into and out of the league | info.playertransferredin // info.playertransferredout // | transfer.in // transfer.out // |
| stats-manager | Handles the storage of the league table and statistics |  | info.resultcompleted // leaguetable.updatedata // leaguetable.updatestats // |

