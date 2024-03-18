package com.java.techstuff.examples.graphs;

import lombok.extern.slf4j.Slf4j;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

@Slf4j
public class BFS {
  private final Queue<String> queue;
  private final List<String> discoveredWebSites;

  public BFS() {
    this.queue = new LinkedList<>();
    this.discoveredWebSites = new ArrayList<>();
  }

  public void discoverWeb(String root) {
    queue.add(root);
    discoveredWebSites.add(root);

    while (!queue.isEmpty()) {
      String vertex = queue.remove();
      String rawHtml = readURL(vertex);

      Document doc = Jsoup.parse(rawHtml);
      Elements links = doc.select("a[href]");

      for (Element link : links) {
        String url = link.absUrl("href");
        if (!discoveredWebSites.contains(url)) {
          discoveredWebSites.add(url);
          queue.add(url);
        }
      }
    }
  }

  private String readURL(String vertex) {
    HttpClient client = HttpClient.newHttpClient();
    HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create(vertex))
            .build();
    client.close();

    try {
      HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
      return response.body();
    } catch (InterruptedException ie) {
      Thread.currentThread().interrupt();
      log.error("Thread interrupted while crawling the website: " + vertex, ie);
      return "";
    } catch (Exception exception) {
      log.error("Errors occur while crawling the website: " + vertex, exception);
      return "";
    }
  }
}
