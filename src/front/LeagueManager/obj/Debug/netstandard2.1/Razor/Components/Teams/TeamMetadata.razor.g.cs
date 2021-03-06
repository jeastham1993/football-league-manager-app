#pragma checksum "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor" "{ff1816ec-aa5e-4d10-87f7-6f4963833460}" "9fc424969de6020a51c387ad82e47a5e1cd4f6be"
// <auto-generated/>
#pragma warning disable 1591
namespace LeagueManager.Components.Teams
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Components;
#nullable restore
#line 1 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using System.Net.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 2 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using Microsoft.AspNetCore.Components.Forms;

#line default
#line hidden
#nullable disable
#nullable restore
#line 3 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using Microsoft.AspNetCore.Components.Routing;

#line default
#line hidden
#nullable disable
#nullable restore
#line 4 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using Microsoft.AspNetCore.Components.Web;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using Microsoft.JSInterop;

#line default
#line hidden
#nullable disable
#nullable restore
#line 6 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using LeagueManager;

#line default
#line hidden
#nullable disable
#nullable restore
#line 7 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using LeagueManager.Shared;

#line default
#line hidden
#nullable disable
#nullable restore
#line 8 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using LeagueManager.Models;

#line default
#line hidden
#nullable disable
#nullable restore
#line 9 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using LeagueManager.Components;

#line default
#line hidden
#nullable disable
#nullable restore
#line 10 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\_Imports.razor"
using LeagueManager.Services;

#line default
#line hidden
#nullable disable
    public partial class TeamMetadata : Microsoft.AspNetCore.Components.ComponentBase
    {
        #pragma warning disable 1998
        protected override void BuildRenderTree(Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder)
        {
#nullable restore
#line 1 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
 if (TeamData == null)
{

#line default
#line hidden
#nullable disable
            __builder.AddContent(0, "    ");
            __builder.AddMarkupContent(1, "<p>Loading...</p>\r\n");
#nullable restore
#line 4 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
}
else
{

#line default
#line hidden
#nullable disable
            __builder.AddContent(2, "    ");
            __builder.OpenElement(3, "h1");
            __builder.AddContent(4, 
#nullable restore
#line 7 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
         TeamData.Name

#line default
#line hidden
#nullable disable
            );
            __builder.CloseElement();
            __builder.AddMarkupContent(5, "\r\n    ");
            __builder.AddMarkupContent(6, "<h2>Current Players</h2>\r\n    ");
            __builder.OpenComponent<LeagueManager.Components.Teams.CreatePlayer>(7);
            __builder.AddAttribute(8, "OnSave", Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<Microsoft.AspNetCore.Components.EventCallback<LeagueManager.Models.Player>>(Microsoft.AspNetCore.Components.EventCallback.Factory.Create<LeagueManager.Models.Player>(this, 
#nullable restore
#line 9 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
                          OnAddPlayer

#line default
#line hidden
#nullable disable
            )));
            __builder.AddAttribute(9, "TeamId", Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.String>(
#nullable restore
#line 9 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
                                                TeamId

#line default
#line hidden
#nullable disable
            ));
            __builder.CloseComponent();
            __builder.AddMarkupContent(10, "\r\n    ");
            __builder.OpenElement(11, "div");
            __builder.AddMarkupContent(12, "\r\n        ");
            __builder.OpenComponent<LeagueManager.Components.Teams.PlayerList>(13);
            __builder.AddAttribute(14, "Players", Microsoft.AspNetCore.Components.CompilerServices.RuntimeHelpers.TypeCheck<System.Collections.Generic.IEnumerable<LeagueManager.Models.Player>>(
#nullable restore
#line 11 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
                              TeamData.Players

#line default
#line hidden
#nullable disable
            ));
            __builder.CloseComponent();
            __builder.AddMarkupContent(15, "\r\n    ");
            __builder.CloseElement();
            __builder.AddMarkupContent(16, "\r\n");
#nullable restore
#line 13 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
}

#line default
#line hidden
#nullable disable
        }
        #pragma warning restore 1998
#nullable restore
#line 15 "C:\Users\jeast\source\Github\league-manager\football-league-manager-app\src\front\LeagueManager\Components\Teams\TeamMetadata.razor"
       
    [Parameter]
    public Team TeamData { get; set; }

    [Parameter]
    public string TeamId { get; set; }

    [Parameter]
    public EventCallback<Player> OnAddPlayer { get; set; }

#line default
#line hidden
#nullable disable
    }
}
#pragma warning restore 1591
