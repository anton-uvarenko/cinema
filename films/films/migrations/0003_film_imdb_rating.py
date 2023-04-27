# Generated by Django 4.1.7 on 2023-03-21 15:54

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('films', '0002_rename_realese_date_film_release_date'),
    ]

    operations = [
        migrations.AddField(
            model_name='film',
            name='imdb_rating',
            field=models.DecimalField(decimal_places=2, default=0.0, max_digits=4),
            preserve_default=False,
        ),
    ]
