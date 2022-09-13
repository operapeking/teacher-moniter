using Microsoft.Toolkit.Uwp.Notifications;
using Newtonsoft.Json;

class SendToast
{
    struct Status
    {
        public DateTime time;
        public string ip;
        public bool isComing;
    }

    static void Send(string title, string content)
    {
        new ToastContentBuilder()
            .AddText(title)
            .AddText(content)
            .Show();
    }
    static readonly HttpClient client = new();
    static async Task Main(string[] args)
    {
        if (args.Length == 0)
        {
            Console.WriteLine("No server provided.");
            return;
        }
        string url = args[0] + "/query";
        int gone = 0;
        for (; ; )
        {
            try
            {
                string responseBody = await client.GetStringAsync(url);
                var result = JsonConvert.DeserializeObject<Status>(responseBody);
                if (result.isComing && gone != 2)
                {
                    Send("Dangerous", "Someone is coming! " + result.time.ToShortTimeString() + " from " + result.ip);
                    gone = 2;
                }
                else if (!result.isComing && gone != 1)
                {
                    Send("Safe", "No one is around. " + result.time.ToShortTimeString() + " from " + result.ip);
                    gone = 1;
                }
                Thread.Sleep(100);
            }
            catch (HttpRequestException e)
            {
                Console.WriteLine("\nException Caught!");
                Console.WriteLine("Message :{0} ", e.Message);
            }
        }
    }
}