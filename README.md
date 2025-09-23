TP 2: Système de Notifications et Logging
Objectif du Projet
Construire une application en ligne de commande qui simule l'envoi de notifications via différents canaux et qui archive un historique de tous les envois réussis.

Fonctionnalités Attendues
Un Système de Notification Flexible

Votre programme doit pouvoir gérer plusieurs types de notifications : Email, SMS, et Push.

Le notificateur SMS doit inclure une logique de validation. Si un numéro de téléphone est considéré comme invalide (par exemple, il ne commence pas par "06"), l'envoi doit échouer et retourner une erreur explicite.
Le programme principal doit pouvoir traiter une liste contenant différents types de notificateurs les uns à la suite des autres.
Les erreurs d'envoi doivent être capturées et affichées dans la console sans stopper le reste du programme.

Un Système d'Archivage (Storer)

Chaque notification envoyée avec succès doit être enregistrée.

L'enregistrement doit contenir le message envoyé ainsi que la date et l'heure de l'envoi (timestamp).
Les enregistrements seront conservés en mémoire.
À la fin de l'exécution, le programme devra afficher la liste complète de toutes les notifications qui ont été archivées.

Concepts Clés à Utiliser : interfaces, structs, slices, gestion d'erreurs, le package time et tout ce qu'on a vu en cours ensemble.