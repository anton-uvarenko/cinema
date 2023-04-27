# Generated by Django 4.1.7 on 2023-04-15 13:14

import actors.validators
from django.db import migrations, models
import films.validators


class Migration(migrations.Migration):

    dependencies = [
        ('actors', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='actor',
            name='age',
            field=models.IntegerField(validators=[actors.validators.validate_age]),
        ),
        migrations.AlterField(
            model_name='actor',
            name='description',
            field=models.TextField(validators=[films.validators.validate_text]),
        ),
        migrations.AlterField(
            model_name='actor',
            name='name',
            field=models.CharField(max_length=50, validators=[films.validators.validate_names]),
        ),
    ]
