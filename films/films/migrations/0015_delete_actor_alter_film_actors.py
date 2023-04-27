# Generated by Django 4.1.7 on 2023-04-14 11:02

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('actors', '0001_initial'),
        ('films', '0014_actor_remove_film_genre_film_genres_film_actors'),
    ]

    operations = [
        migrations.DeleteModel(
            name='Actor',
        ),
        migrations.AlterField(
            model_name='film',
            name='actors',
            field=models.ManyToManyField(db_table='FilmActor', to='actors.actor'),
        ),
    ]
